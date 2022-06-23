# mlflow-log-cli

## Installation

```sh
pip install mlflow-log-cli
```

## Usage

```sh
# Set MLflow tracking URI
export MLFLOW_TRACKING_URI=http://127.0.0.1:5000

# Set MLflow experiment name or ID
export MLFLOW_EXPERIMENT_NAME=
or
export MLFLOW_EXPERIMENT_ID=

# log metrics, params and tags from json file
mlflow-log-cli -f template.json
```

## Templates

[template.json](template.json)