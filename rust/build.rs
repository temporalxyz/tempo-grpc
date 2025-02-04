fn main() {
    let includes: &[&str] = &["../protos/"];
    tonic_build::configure()
        .bytes(&["tempo.Transaction.payload"])
        .compile_protos(&["../protos/tempo.proto"], includes)
        .unwrap_or_else(|e| panic!("Failed to compile protos {:?}", e));
}
