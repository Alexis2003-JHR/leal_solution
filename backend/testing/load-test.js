import http from 'k6/http';
import { sleep, check } from 'k6';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';


function randomFloatBetween(min, max) {
    return Math.random() * (max - min) + min;
}

export let options = {
    stages: [
        { duration: '15s', target: 150 },
        { duration: '30s', target: 300 },
        { duration: '15s', target: 150 }, 
    ],
};

export default function () {
    let payload = JSON.stringify({
        usuario: {
            numero_documento: 1007105219
        },
        id_sucursal: 10,
        valor: Number(randomFloatBetween(5000, 200000).toFixed(2))
    });

    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    let res = http.post('http://localhost:6060/api/v1/transactions', payload, params);

    check(res, {
        'status is 200': (r) => r.status === 200,
        'response time < 500ms': (r) => r.timings.duration < 500,
    });

    sleep(1);
}
