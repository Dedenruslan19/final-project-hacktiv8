import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate, Trend, Counter } from 'k6/metrics';

// Custom metrics
const errorRate = new Rate('errors');
const bidDuration = new Trend('bid_duration');
const successfulBids = new Counter('successful_bids');
const failedBids = new Counter('failed_bids');

// STRESS TEST - Push system to its limits
export const options = {
  stages: [
    { duration: '1m', target: 100 },    // Warm up
    { duration: '2m', target: 500 },    // Ramp to 500 users
    { duration: '3m', target: 1000 },   // Stress: 1000 concurrent users
    { duration: '2m', target: 1500 },   // Peak stress: 1500 users
    { duration: '3m', target: 1500 },   // Hold at peak
    { duration: '2m', target: 500 },    // Recover
    { duration: '1m', target: 0 },      // Cool down
  ],
  thresholds: {
    'http_req_duration': ['p(95)<5000'],  // Relaxed threshold for stress test
    'http_req_failed': ['rate<0.2'],      // Allow 20% error rate under stress
    'errors': ['rate<0.2'],
  },
};

const BASE_URL = __ENV.BASE_URL || 'https://your-gcp-url.run.app';
const SESSION_ID = '1';  // Always use session 1 for testing
const ITEM_ID = '1';     // Always use item 1 for testing
const AUTH_TOKEN = __ENV.AUTH_TOKEN || 'your-jwt-token-here';

export default function () {
  const bidAmount = 100000 + Math.floor(Math.random() * 900000);
  const url = `${BASE_URL}/auction/sessions/${SESSION_ID}/items/${ITEM_ID}/bid`;
  
  const payload = JSON.stringify({
    amount: bidAmount,
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${AUTH_TOKEN}`,
    },
    timeout: '10s', // Increase timeout for stress test
  };

  const response = http.post(url, payload, params);
  
  bidDuration.add(response.timings.duration);

  const success = check(response, {
    'status is 2xx or 409': (r) => (r.status >= 200 && r.status < 300) || r.status === 409,
    'not timeout': (r) => r.status !== 0,
  });

  if (response.status === 200 || response.status === 201) {
    successfulBids.add(1);
  } else if (response.status !== 409 && response.status !== 0) {
    failedBids.add(1);
    errorRate.add(1);
  }

  // Minimal think time for stress test
  sleep(0.5);
}

export function setup() {
  console.log('=== STRESS TEST START ===');
  console.log(`Target: ${BASE_URL}`);
  console.log(`Peak Load: 1500 concurrent users`);
  console.log('=========================');
}
