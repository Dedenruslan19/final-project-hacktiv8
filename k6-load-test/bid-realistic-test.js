import http from 'k6/http';
import { check, sleep } from 'k6';
import { Counter, Rate, Trend } from 'k6/metrics';

// Custom metrics
const successfulBids = new Counter('successful_bids');
const conflictBids = new Counter('conflict_bids_409');
const tooLowBids = new Counter('too_low_bids');
const duplicateBids = new Counter('duplicate_bids');
const errorRate = new Rate('errors');
const bidLatency = new Trend('bid_latency');

// REALISTIC AUCTION SCENARIO
// Simulates real auction behavior with multiple bidders competing
export const options = {
  scenarios: {
    // Scenario 1: Active bidders (aggressive)
    active_bidders: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '30s', target: 50 },
        { duration: '3m', target: 100 },
        { duration: '2m', target: 150 },
        { duration: '1m', target: 50 },
        { duration: '30s', target: 0 },
      ],
      gracefulRampDown: '30s',
      exec: 'activeBidder',
    },
    
    // Scenario 2: Casual observers (check prices occasionally)
    casual_observers: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '1m', target: 200 },
        { duration: '4m', target: 300 },
        { duration: '2m', target: 100 },
      ],
      gracefulRampDown: '30s',
      exec: 'casualObserver',
    },

    // Scenario 3: Last minute rush
    last_minute_rush: {
      executor: 'ramping-arrival-rate',
      startRate: 1,
      timeUnit: '1s',
      preAllocatedVUs: 50,
      maxVUs: 500,
      stages: [
        { duration: '5m', target: 10 },   // Steady
        { duration: '1m', target: 100 },  // Rush starts
        { duration: '30s', target: 200 }, // Peak rush
        { duration: '30s', target: 5 },   // Cool down
      ],
      exec: 'lastMinuteBidder',
      startTime: '4m',
    },
  },
  thresholds: {
    'http_req_duration': ['p(95)<2000', 'p(99)<3000'],
    'http_req_failed': ['rate<0.15'],
    'bid_latency': ['p(95)<1500', 'p(99)<2500'],
  },
};

const BASE_URL = __ENV.BASE_URL || 'https://your-gcp-url.run.app';
const SESSION_ID = '1';  // Always use session 1 for testing
const ITEM_ID = '1';     // Always use item 1 for testing
const AUTH_TOKEN = __ENV.AUTH_TOKEN || 'your-jwt-token-here';

// Shared state for highest bid tracking (approximate)
let approximateHighestBid = 100000;

// Active bidder: Places bids frequently with incremental amounts
export function activeBidder() {
  // Get current highest + increment
  const increment = Math.floor(Math.random() * 3) * 10000 + 10000; // 10k, 20k, or 30k
  const bidAmount = approximateHighestBid + increment;
  
  const result = placeBid(bidAmount, 'active');
  
  if (result.status === 200 || result.status === 201) {
    approximateHighestBid = bidAmount;
  }
  
  // Active bidders wait 5-15 seconds between bids
  sleep(5 + Math.random() * 10);
}

// Casual observer: Mostly watches, occasionally bids
export function casualObserver() {
  // 80% chance to just check status, 20% chance to bid
  if (Math.random() < 0.2) {
    const bidAmount = approximateHighestBid + Math.floor(Math.random() * 50000) + 20000;
    placeBid(bidAmount, 'casual');
  }
  
  // Casual users check less frequently
  sleep(10 + Math.random() * 20);
}

// Last minute bidder: Aggressive bidding near the end
export function lastMinuteBidder() {
  // Last minute bidders place higher bids
  const aggressiveIncrement = Math.floor(Math.random() * 5) * 10000 + 50000; // 50k-90k
  const bidAmount = approximateHighestBid + aggressiveIncrement;
  
  const result = placeBid(bidAmount, 'lastminute');
  
  if (result.status === 200 || result.status === 201) {
    approximateHighestBid = bidAmount;
  }
  
  // Very short wait time - competing aggressively
  sleep(1 + Math.random() * 3);
}

function placeBid(amount, bidderType) {
  const url = `${BASE_URL}/auction/sessions/${SESSION_ID}/items/${ITEM_ID}/bid`;
  
  const payload = JSON.stringify({
    amount: amount,
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${AUTH_TOKEN}`,
    },
    tags: { 
      name: 'PlaceBid',
      bidderType: bidderType,
    },
  };

  const startTime = new Date().getTime();
  const response = http.post(url, payload, params);
  const endTime = new Date().getTime();
  
  bidLatency.add(endTime - startTime);

  check(response, {
    'status is 2xx': (r) => r.status >= 200 && r.status < 300,
    'response time OK': (r) => r.timings.duration < 3000,
  });

  // Track different response types
  if (response.status === 200 || response.status === 201) {
    successfulBids.add(1);
  } else if (response.status === 409) {
    conflictBids.add(1);
    const body = response.body.toLowerCase();
    if (body.includes('too low')) {
      tooLowBids.add(1);
    } else if (body.includes('duplicate')) {
      duplicateBids.add(1);
    }
  } else {
    errorRate.add(1);
    console.log(`[${bidderType}] Error: ${response.status} - ${response.body}`);
  }

  return {
    status: response.status,
    duration: response.timings.duration,
  };
}

export function setup() {
  console.log('=== REALISTIC AUCTION SCENARIO ===');
  console.log(`Base URL: ${BASE_URL}`);
  console.log(`Session ID: ${SESSION_ID}, Item ID: ${ITEM_ID}`);
  console.log('Scenarios:');
  console.log('  1. Active Bidders: 50-150 concurrent users');
  console.log('  2. Casual Observers: 100-300 concurrent users');
  console.log('  3. Last Minute Rush: Up to 200 req/s spike');
  console.log('===================================');
  
  return { startTime: new Date().toISOString() };
}

export function teardown(data) {
  console.log('=== TEST COMPLETED ===');
  console.log(`Started: ${data.startTime}`);
  console.log(`Ended: ${new Date().toISOString()}`);
  console.log('======================');
}
