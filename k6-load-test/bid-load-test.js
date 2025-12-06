import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate, Trend, Counter } from 'k6/metrics';

// Custom metrics
const errorRate = new Rate('errors');
const bidDuration = new Trend('bid_duration');
const successfulBids = new Counter('successful_bids');
const failedBids = new Counter('failed_bids');

// Test configuration
export const options = {
  stages: [
    { duration: '30s', target: 50 },   // Ramp up to 50 users
    { duration: '1m', target: 100 },   // Ramp up to 100 users
    { duration: '2m', target: 200 },   // Ramp up to 200 users
    { duration: '3m', target: 500 },   // Peak load: 500 concurrent users
    { duration: '2m', target: 200 },   // Ramp down to 200 users
    { duration: '1m', target: 100 },   // Ramp down to 100 users
    { duration: '30s', target: 0 },    // Ramp down to 0 users
  ],
  thresholds: {
    'http_req_duration': ['p(95)<2000', 'p(99)<3000'], // 95% requests < 2s, 99% < 3s
    'http_req_failed': ['rate<0.1'],   // Error rate < 10%
    'errors': ['rate<0.1'],            // Custom error rate < 10%
    'bid_duration': ['p(95)<1500'],    // 95% bid requests < 1.5s
  },
};

// Configuration - EDIT THESE VALUES
const BASE_URL = __ENV.BASE_URL || 'https://your-gcp-url.run.app';
const SESSION_ID = '1';  // Always use session 1 for testing
const ITEM_ID = '1';     // Always use item 1 for testing
const AUTH_TOKEN = __ENV.AUTH_TOKEN || 'your-jwt-token-here';

// Generate random user IDs for simulation (User 1-100)
const USER_COUNT = 100;
const userTokens = generateUserTokens(USER_COUNT);

function generateUserTokens(count) {
  // In real scenario, you should have actual JWT tokens for different users
  // For now, we'll use the same token but simulate different users
  const tokens = [];
  for (let i = 1; i <= count; i++) {
    tokens.push({
      userId: i,
      token: AUTH_TOKEN, // In production, each user should have different token
    });
  }
  return tokens;
}

// Shared counter for incremental bidding (simulates real auction behavior)
let bidCounter = 0;

export default function () {
  // Randomly select a user
  const userIndex = Math.floor(Math.random() * userTokens.length);
  const user = userTokens[userIndex];

  // Incremental bidding: each bid is higher than the previous
  // Start from 250,000 (starting_price) and increment by 10,000-50,000 each time
  bidCounter++;
  const incrementPerBid = 10000 + Math.floor(Math.random() * 40000); // 10k-50k increment
  const bidAmount = 250000 + (bidCounter * incrementPerBid);

  const url = `${BASE_URL}/auction/sessions/${SESSION_ID}/items/${ITEM_ID}/bid`;
  
  const payload = JSON.stringify({
    amount: bidAmount,
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${user.token}`,
    },
    tags: { 
      name: 'PlaceBid',
      userId: user.userId,
    },
  };

  // Send bid request
  const response = http.post(url, payload, params);

  // Record custom metrics
  bidDuration.add(response.timings.duration);

  // Check response
  const success = check(response, {
    'status is 200 or 201': (r) => r.status === 200 || r.status === 201,
    'status is 409 (conflict - expected for duplicate/low bids)': (r) => r.status === 409,
    'response has body': (r) => r.body.length > 0,
    'response time < 3000ms': (r) => r.timings.duration < 3000,
  });

  // Count success/failure
  if (response.status === 200 || response.status === 201) {
    successfulBids.add(1);
  } else if (response.status !== 409) {
    // 409 is expected (bid too low, duplicate, etc), don't count as failure
    failedBids.add(1);
    errorRate.add(1);
    console.log(`Error: ${response.status} - ${response.body}`);
  } else {
    // 409 is business logic rejection, not system error
    errorRate.add(0);
  }

  // Log sample responses for debugging
  if (__ITER % 100 === 0) {
    console.log(`[User ${user.userId}] Bid: ${bidAmount}, Status: ${response.status}, Duration: ${response.timings.duration}ms`);
  }

  // Think time: random delay between 1-3 seconds
  sleep(Math.random() * 2 + 1);
}

// Setup function runs once per VU before the default function
export function setup() {
  console.log('=== Load Test Setup ===');
  console.log(`Base URL: ${BASE_URL}`);
  console.log(`Session ID: ${SESSION_ID}`);
  console.log(`Item ID: ${ITEM_ID}`);
  console.log(`Simulated Users: ${USER_COUNT}`);
  console.log('=======================');
}

// Teardown function runs once after all iterations complete
export function teardown(data) {
  console.log('=== Load Test Complete ===');
  console.log('Check the results above for metrics');
  console.log('==========================');
}
