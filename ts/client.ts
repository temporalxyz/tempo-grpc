import * as grpc from '@grpc/grpc-js';
import { TransactionStreamClient, StartStream, Transaction } from './generated/tempo'; // Adjust path if needed

function main() {
    // Replace with the actual endpoint.
    const endpoint = 'check the docs for endpoints';

    // Create the gRPC client using the generated service client.
    const client = new TransactionStreamClient(
        endpoint,
        grpc.credentials.createInsecure()
    );

    // Build the StartStream request.
    const request: StartStream = {
        authToken: "dm JC, cavey, ben, or j for auth token", // Replace with your auth token.
    };

    // Open the transaction stream.
    const stream = client.openTransactionStream(request);

    // Listen for stream events.
    stream.on('data', (transaction: Transaction) => {
        const { slot, index, payload } = transaction;
        snipe(slot, index, payload);
    });

    stream.on('end', () => {
        console.log('TransactionStream closed!');
    });

    stream.on('error', (err: any) => {
        console.error('Stream error:', err);
    });
}

// A placeholder function to process each transaction.
function snipe(slot: number, index: number, _payload: Uint8Array) {
    // console.log(`Snipe: slot=${slot}, index=${index}, payload=${Buffer.from(_payload).toString('hex')}`);
    console.log(`tsSnipe: slot=${slot}, index=${index}`);
}

main();
