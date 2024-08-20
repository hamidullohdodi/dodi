CREATE TABLE product_categories (
        id UUID PRIMARY KEY default gen_random_uuid(),
        name VARCHAR(50) NOT NULL,
        description TEXT,
        created_at varchar(120),
        updated_at varchar(120),
        deleted_at  varchar(120)
);


CREATE TABLE products (
        id UUID PRIMARY KEY default gen_random_uuid(),
        name VARCHAR(100) NOT NULL,
        description TEXT,
        price DECIMAL(10, 2) NOT NULL,
        category_id UUID REFERENCES product_categories(id),
        artisan_id varchar(120),
        quantity INTEGER NOT NULL,
        created_at varchar(120),
        updated_at varchar(120),
        deleted_at  bigint default 0
);


CREATE TABLE orders (
        id UUID PRIMARY KEY default gen_random_uuid(),
        user_id uuid not null,
        total_amount DECIMAL(10, 2) NOT NULL,
        status VARCHAR(20) NOT NULL,
        shipping_address JSONB NOT NULL,
        created_at varchar(120),
        updated_at varchar(120),
        deleted_at  varchar(120)
);

CREATE TABLE order_items (
        id UUID PRIMARY KEY default gen_random_uuid(),
        order_id UUID REFERENCES orders(id),
        product_id UUID REFERENCES products(id),
        quantity INTEGER NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        created_at varchar(120),
        updated_at varchar(120),
        deleted_at  varchar(120)
);

CREATE TABLE ratings (
        id UUID PRIMARY KEY default gen_random_uuid(),
        product_id UUID REFERENCES products(id),
        user_id uuid not null ,
        rating DECIMAL(2, 1) NOT NULL,
        comment TEXT,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments (
        id UUID PRIMARY KEY default gen_random_uuid(),
        order_id UUID REFERENCES orders(id),
        amount DECIMAL(10, 2),
        status VARCHAR(20),
        transaction_id VARCHAR(100),
        payment_method VARCHAR(50),
        created_at varchar(120),
        updated_at varchar(120),
        deleted_at  varchar(120)
);

CREATE TABLE shipping_info (
                               id UUID PRIMARY KEY default gen_random_uuid(),
                               order_id UUID NOT NULL REFERENCES orders(id),
                               tracking_number VARCHAR(50),
                               carrier VARCHAR(50),
                               estimated_delivery_date VARCHAR(120),
                               created_at VARCHAR(120),
                               updated_at VARCHAR(120),
                               deleted_at VARCHAR(120)
);

CREATE TABLE order_payments (
                                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                order_id UUID NOT NULL REFERENCES orders(id),
                                payment_method VARCHAR(50) NOT NULL,
                                card_number VARCHAR(20) NOT NULL,
                                expiry_date VARCHAR(10) NOT NULL,
                                cvv VARCHAR(4) NOT NULL,
                                created_at VARCHAR(120),
                                updated_at VARCHAR(120),
                                deleted_at VARCHAR(120)
);
