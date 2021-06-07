# Auth-app
This API serves to generate JWT tokens which will later be used by the fetch-app to get commodity data. The features available in this API are register, login, and profile. Using a simple architecture but able to handle authentication well. You can find complete documentation regarding this application at the following [link](https://github.com/hudabikhoir/efishery-monorepo/blob/master/API.md)

***Available role:***
| ID | Role |
| --- | ------|
| 1 | Admin |
| 2 | Finance |
| 3 | Developer |

## Technology
- Python 3.8
- SQLite
- Flask
- Flask Migrate
- SQLAlchemy

## Installation
- `virtualenv -p python3` venv if you don't have virtualenv you can install with `pip3 install virtualenv`
- `source env/bin/activate` to activate your venv
- open your terminal and then export this variable environment
```
    $ export FLASK_DEBUG=1
    $ export FLASK_APP=app.py
```
- intall all package dependecies with `pip install requirements.txt`
- preparing data migration
```
    $ flask db init
    $ flask db migrate
    $ flask db upgrade
```
-  now run your auth-app `flask run`
- you can access `http://127.0.0.1:5000/` on your browser or postman