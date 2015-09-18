  ##########
  # Database
  ##########
  # Database config all start with db_ and can be set either in the config
  # file or with an environment variable (DB_NAME for example).

  config.db_driver = 'mongo'
  config.db_name = (config.app_name + '_' + Volt.env.to_s)
  config.db_host = 'mongo'
  config.db_port = 27017
