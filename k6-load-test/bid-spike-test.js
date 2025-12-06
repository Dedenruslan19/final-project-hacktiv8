import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate, Counter, Trend } from 'k6/metrics';

// Custom metrics
const successfulBids = new Counter('successful_bids');
const failedBids = new Counter('failed_bids');
const errorRate = new Rate('errors');
const bidDuration = new Trend('bid_duration');

// LOAD TEST - 50 VUs for 30 seconds
export const options = {
  vus: 50,
  duration: '30s',
  thresholds: {
    'http_req_duration': ['p(95)<2000', 'p(99)<3000'],
    'http_req_failed': ['rate<0.1'],
    'errors': ['rate<0.1'],
    'bid_duration': ['p(95)<1500'],
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8000';
const SESSION_ID = __ENV.SESSION_ID || '1';
const ITEM_ID = __ENV.ITEM_ID || '1';
const AUTH_TOKEN = __ENV.AUTH_TOKEN || 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRlZGVuM0BtYWlsLmNvbSIsImV4cCI6MTc2NTAxNzA3NywiaWQiOjIwLCJyb2xlIjoidXNlciJ9.tWImuhepBUnMRBdMl9WzeceV1brX7dUxJ9KrT_uixXc';

// Shared counter for incremental bidding
let bidCounter = 0;

export default function () {
  // Incremental bidding: start from 250k and increment by 10k-50k
  bidCounter++;
  const increment = 10000 + Math.floor(Math.random() * 40000); // 10k-50k
  const bidAmount = 250000 + (bidCounter * increment);
  
  const url = `${BASE_URL}/auction/sessions/${SESSION_ID}/items/${ITEM_ID}/bid`;
  
  const payload = JSON.stringify({
    amount: bidAmount,
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${AUTH_TOKEN}`,
    },
    timeout: '10s',
  };

  const startTime = new Date().getTime();
  const response = http.post(url, payload, params);
  const endTime = new Date().getTime();
  
  bidDuration.add(endTime - startTime);

  check(response, {
    'status is 200 or 201': (r) => r.status === 200 || r.status === 201,
    'status is 409 (conflict - expected)': (r) => r.status === 409,
    'response has body': (r) => r.body && r.body.length > 0,
    'response time < 3000ms': (r) => r.timings.duration < 3000,
  });

  if (response.status === 200 || response.status === 201) {
    successfulBids.add(1);
    if (__ITER % 10 === 0) {
      console.log(`[VU ${__VU}] SUCCESS - Bid: Rp ${bidAmount.toLocaleString()}, Duration: ${response.timings.duration.toFixed(0)}ms`);
    }
  } else if (response.status === 409) {
    // 409 is normal for auction competition (bid too low or duplicate)
  } else {
    failedBids.add(1);
    errorRate.add(1);
    if (__ITER % 5 === 0) {
      console.log(`[VU ${__VU}] ERROR ${response.status} - Bid: Rp ${bidAmount.toLocaleString()}`);
      console.log(`Response: ${response.body.substring(0, 100)}`);
    }
  }

  // Small delay to reduce load
  sleep(0.5);
}

export function setup() {
  console.log('=== LOAD TEST START ===');
  console.log(`Base URL: ${BASE_URL}`);
  console.log(`VUs: 50, Duration: 30s`);
  console.log(`Session ID: ${SESSION_ID}, Item ID: ${ITEM_ID}`);
  console.log('Bid Strategy: Incremental (250k + counter * 10k-50k)');
  console.log('=======================');
}

export function teardown() {
  console.log('=== LOAD TEST COMPLETE ===');
  console.log('Check results above for details');
  console.log('==========================');
}
