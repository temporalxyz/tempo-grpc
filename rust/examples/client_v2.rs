use tempo_protos::{transaction_stream_client::TransactionStreamClient, StartStreamV2, Transaction};

#[tokio::main(flavor = "current_thread")]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Connect to endpoint
    let endpoint = "check the docs for endpoints";

    let mut client = TransactionStreamClient::connect(endpoint).await?;

    // Send request to start stream
    let auth_token = "dm JC, cavey, ben, or j for auth token".to_string(); /* TODO */
    let start_stream_request = StartStreamV2 { auth_token, static_account_filter: vec![
        /* TODO: your STATIC accounts (programs, signers, writable, non-LUT readable) */
    ], };
    let mut stream = client.open_transaction_stream_v2(start_stream_request).await?;

    // Process transactions from stream
    loop {
        match stream.get_mut().message().await {
            // Next message successfully received
            Ok(Some(Transaction {
                // Slot for the shred this transaction was found in
                slot,
                // The transaction's index within the block in this slot
                index,
                // The transaction bytes
                payload,
            })) => {
                let snipe = || /* TODO */ 
                println!("received tx {index} in slot {slot}");
                (slot, index, payload);
                snipe();
            }

            // Stream was closed. Should not happen unless server goes down.
            // May want some stream reconnect logic
            Ok(None) => {
                println!("TransactionStream closed!");
                return Ok(());
            }

            // Invalid Message. Also should not happen. Maybe yell at us
            Err(e) => {
                println!("invalid stream message! error is {e:?}");
                return Err(e.into());
            }
        }
    }
}
