# Guppy

Little fish-like gRPC services with gRPC server relfection containerised and ready to use as

    docker run -p9090:9090 julia/echo

Alternatively, clone, build and run locally with

    make run-echo

Test with

    grpcurl -plaintext localhost:9090 describe
    grpcurl -plaintext -d '{"message" : "👋"}'  localhost:9090 echo.Echo/Hello

For the demo Routeguide service try `docker run -p9090:9090 julia/routeguide` or `make run-rguide` locally.

    
