db.createUser(
    {
        user: "mongo",
        pwd: "authpass",
        roles: [
            {
                role: "readWrite",
                db: "auth"
            }
        ]
    }
);