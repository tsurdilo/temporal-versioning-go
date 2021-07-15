## Temporal Versioning Demo - Go SDK

This example demonstrates how to make and version changes to your workflow using Temporal Go SDK.

The sample is a simple customer rewards workflow. In this demo we will make 2 changes to it while the workflow
is running.

### Getting started

#### 1. Start the Temporal Server:
First we need to start the Temporal Server on local Docker.
Note that in order to be able to query workflows by their change id and version in the web-ui 
or via tctl we have to start the Temporal Server with Elasticsearch enabled:

```shell script
git clone https://github.com/temporalio/docker-compose.git
cd  docker-compose
docker-compose -f docker-compose-cas-es.yml up
```

In case you don't want to run the queries agaist workflow changes, you can start Temporal Server on Docker
without Elasticsearch enabled:

```shell script
git clone https://github.com/temporalio/docker-compose.git
cd  docker-compose
docker-compose -f docker-compose-cas-es.yml up
```

#### 2. Run the demo

a) Start the workflow worker

```shell script
go run worker/main.go
```

b) Start the workflow starter

```shell script
go run starter/main.go
```

Open the Temporal Web-UI on [localhost:8080](http://localhost:8080) and see your workflow is running.

c) Let's make our first change. 

Note that for both of the changes you have about 2 minutes to do all these stops. If you want to increase this 
amount of time change the "DemoWaitDuration" amount in the workflow starter.

Go to our workflow code "workflow/workflow.go" and uncomment the code between:

```shell script
// CHANGE 1
< UNCOMMENT ALL CODE HERE >
// END CHANGE 1
```

Now let's test that our change will not break the determinism of any already running workflow instances.

Open "tests/replay_test.go" and run the "TestReplayFromInitialVersion" test method. This will do a replay against 
the workflow events of our workflow execution before we applied the changes.

Next step is to restart our worker, stop the running workflow worker and start it again. This will apply 
our changes to our running workflow instances.

d) Query the workflow execution by its change

If you started Temporal Server with Elasticsearch enabled you can via tctl or the Web-ui run a query to make sure our
workflow has been tagged with the made chage:

```shell script
tctl workflow list --query='TemporalChangeVersion="addedCheck-1"
```

e) Make our second change:

In our "workflow/workflow.go" file now uncomment all the code (note its in two places!) 
between the comments:

```shell script
// CHANGE 2
< UNCOMMENT ALL CODE HERE >
// END CHANGE 2
```

Now let's again stop and start our workflow worker.

You can run the query again to see that the workflow instance has been tagged also with our new workflow changes:

```shell script
tctl workflow list --query='TemporalChangeVersion="addedBonus-1"
```

Wait until workflow execution ends. You should see the result of "200" dollars bonus that has been added
to the customer account.
