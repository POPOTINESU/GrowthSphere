schema "public" {
    name = "public"
}

enum "auth_providers" {
    schema = schema.public
    values = ["local"]
}

enum "permissions" {
    schema = schema.public
    values = ["user", "admin"]
}

table "users" {
    schema = schema.public

    column "id" {
        type = uuid
        null = false
        default = sql("gen_random_uuid()")
    }

    column "username" {
        type = text
        null = false
        default = "anonymous"
    }

    column "email" {
        type = text
        null = false
    }

    column "password" {
        type = text
        null = true
    }

    column "provider" {
        type = text
        null = false
        default = "local"
    }

    column "permissions" {
        type = text
        null = false
        default = "user"
    }

    primary_key {
        columns = [column.id]
    }

    index "unique_email" {
        unique = true
        columns = [column.email]
    }

    index "unique_username" {
        unique = true
        columns = [column.username]
    }

    check "username_length" {
        expr = "(length(username) > 3)"
    }

    check "password_length" {
        expr = "(length(password) > 8)"
    }
}