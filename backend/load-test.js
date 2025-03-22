import http from 'k6/http';
import { sleep, check } from 'k6';
 
 
function randomFloatBetween(min, max) {
    return Math.random() * (max - min) + min;
}

export let options = {
    stages: [
        { duration: '5s', target: 10 },
        { duration: '20s', target: 20 },
        { duration: '5s', target: 10 }, 
    ],
};

export default function () {
    let payload = JSON.stringify({
        usuario: {
            numero_documento: 1007105219
        },
        id_sucursal: 10,
        valor: 5000
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