        CREATE TABLE IF NOT EXISTS users (
        id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
        first_name VARCHAR(50) ,
        email VARCHAR(100) ,
        password VARCHAR(255) ,
        last_name VARCHAR(100),
        date_of_birth DATE,
        role VARCHAR(50) ,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW(),
        deleted_at BIGINT DEFAULT 0
    );
    
    CREATE TABLE IF NOT EXISTS tokens (
        id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
        user_id UUID NOT NULL REFERENCES users(id),
        token VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW(),
        deleted_at BIGINT DEFAULT 0
    );
