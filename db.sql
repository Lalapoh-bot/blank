CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE bookings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL,
    employee_id UUID NOT NULL,
    lapangan_id UUID NOT NULL,
    booking_date DATE NOT NULL,
    jam TIME NOT NULL,
    total_pembayaran NUMERIC(20, 2) NOT NULL,
    pembayaran_sebagian NUMERIC(20, 2) NOT NULL,
    status_pembayaran VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

