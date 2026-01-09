import path from 'path';
import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';

const PROTO_PATH = path.join(process.cwd(), 'proto', 'spotify_curate.proto');

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const fs = require('fs');

const protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
const spotifyCurate = protoDescriptor.proto.SpotifyCurate;

// Load the server certificate
// Note: In production, this should be trusted by the system or passed via env
const certPath = path.join(process.cwd(), '../src/server/certs/server.crt');
const rootCert = fs.readFileSync(certPath);

const client = new spotifyCurate(
    'localhost:50051',
    grpc.credentials.createSsl(rootCert)
);

export const grpcAsync = (method, request) => {
    return new Promise((resolve, reject) => {
        client[method](request, (err, response) => {
            if (err) {
                reject(err);
            } else {
                resolve(response);
            }
        });
    });
};

export default client;
