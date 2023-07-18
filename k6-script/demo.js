import http from 'k6/http';
import {sleep} from 'k6';

export const options = {
  stages: [
    { duration: '1m', target: 10 }, // traffic ramp-up from 1 to 10 users over a minute.
    { duration: '5m', target: 30 }, // stay at 30 users for 5 minutes
    { duration: '1m', target: 0 }, // ramp-down to 0 users
  ],
};

export default function () {
  http.get('https://test.k6.io');
  sleep(1);
}