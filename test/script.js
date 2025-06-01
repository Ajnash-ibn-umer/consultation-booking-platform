import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 10,
  duration: '30s',
  cloud: {
    // Project: Default project
    projectID: 3773583,
    // Test runs with the same name groups test runs together.
    name: 'Test (01/06/2025-13:48:44)'
  }
};

export default function() {
  http.get('http://localhost:8000/v1/auth/role');
  sleep(1);
}