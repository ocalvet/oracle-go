## Setup

1. Create an account in docker hub and accept the oracle terms
2. Install the [instant client](https://www.oracle.com/technetwork/topics/linuxx86-64soft-092277.html)
3. Login to docker hub `$ docker login`
4. Run an oracle container `$ docker run -d -it --name oracle -P store/oracle/database-enterprise:12.2.0.1-slim`
