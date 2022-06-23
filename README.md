
# newrelic-python-paketo-buildpack
## The Paketo New Relic Agent Buildpack is a Cloud Native Buildpack that contributes and configures the New Relic Python Agent.

  

**Behavior**

  
This buildpack will participate if all the following conditions are met

<br/>
The `$BP_NEW_RELIC_ENABLED` is set to true (defaults to false)
<br/><br/>
For Python applications a Procfile is required at root of your application during build stage, sample file available /resources/Procfile
<br/><br/>

The buildpack will do the following for Python applications:
<br/>
Installs New Relic Python agent using pip3
<br/>
Copies New Relic Python Agent configuration file in /resources/newrelic.ini to the root folder of your application
<br/>

**Variables**
<br/>

| Key | Description |
|--|--|
| `BP_NEW_RELIC_ENABLED` | Defaults to false - will not participate - as set in buildpack.toml   |
| `NEW_RELIC_APP_NAME` | Defaults to app_name variable set in newrelic.ini  |
| `NEW_RELIC_LICENSE_KEY`  | Required at build time or runtime     |
| `NEW_RELIC_AGENT_ENABLED`  | Defaults to agent_enabled variable set in newrelic.ini |

<br/>
You can override any setting from a system property or in the newrelic.ini by setting an environment variable.

The environment variable corresponding to a given setting in the config file is the setting name prefixed by NEW_RELIC with all dots (.) and dashes (-) replaced by underscores (_). 

For this to work as part of the **build stage**, you will need to precede the variables with BPE, i.e. `BPE_NEW_RELIC_LOG_LEVEL`

For this to work during **runtime**, simply prefix the setting name, i.e. `NEW_RELIC_LOG_LEVEL`

Please refer Python agent documentation for more information

https://docs.newrelic.com/docs/apm/agents/python-agent/configuration/python-agent-configuration/

<br/>

**Example**
  
pack build CONTAINERNAME -p ./PATHTOPYTHONAPP -b paketo-buildpacks/python -b ./PATHTOLOCALBUILDPACK ... \

--env `BP_NEW_RELIC_ENABLED`=true \ -----------> Required condition to build the buildpack as default is false

--env `BPE_NEW_RELIC_APP_NAME`=xxxxxxxxxx \ -----------> Optional can be set on runtime

--env `BPE_NEW_RELIC_LICENSE_KEY`=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx -----------> Optional can be set on runtime


<br/>

By default the Python agent configuration file will be located at your application root folder, this can be overwritten at runtime, with configmap for kubernetes deployments.


Please refer to Kubernetes documentation for more information about configmaps

[https://kubernetes.io/docs/concepts/configuration/configmap/](https://kubernetes.io/docs/concepts/configuration/configmap/)
