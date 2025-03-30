-- Create the main Budget table
CREATE TABLE budget (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_type VARCHAR(50) NOT NULL, -- contract/engagement/reception
    total_guest INT NOT NULL,
    total_price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

-- Create Venue Table
CREATE TABLE venue (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    venue_name TEXT NOT NULL,
    venue_type VARCHAR(50) NOT NULL, -- hotel/restaurant/other
    location TEXT NOT NULL,
    capacity INT NOT NULL,
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field

);

-- Create Catering Table
CREATE TABLE catering (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    pax INT NOT NULL,
    buffet INT NOT NULL, -- prasmanan
    price INT NOT NULL, -- total price
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

-- Create Menu Table
CREATE TABLE menu (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    catering_id UUID REFERENCES catering(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    item_price INT NOT NULL -- price per pax
);

-- Create Entertainment Table
CREATE TABLE entertainment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    sound_system BOOLEAN NOT NULL DEFAULT FALSE,
    master_ceremony BOOLEAN NOT NULL DEFAULT FALSE,
    duration INT NOT NULL, -- in hours
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

-- Create Performer Table
CREATE TABLE performer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    entertainment_id UUID REFERENCES entertainment(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    type_performer VARCHAR(50) NOT NULL, -- acoustic/electro/other art
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

-- Create Other Service Tables
CREATE TABLE decoration (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    decoration_type VARCHAR(50) NOT NULL, -- floral/lighting/other
    size INT NOT NULL, -- in square meters
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

CREATE TABLE documentation (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    is_fotographer BOOLEAN NOT NULL DEFAULT FALSE,
    is_videographer BOOLEAN NOT NULL DEFAULT FALSE,
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

CREATE TABLE attire (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    is_bride_attire BOOLEAN NOT NULL DEFAULT FALSE,
    is_groom_attire BOOLEAN NOT NULL DEFAULT FALSE,
    is_family_attire BOOLEAN NOT NULL DEFAULT FALSE,
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

CREATE TABLE transportation (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    transportation_type VARCHAR(255) NOT NULL, -- shuttle/car/other
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

CREATE TABLE makeup_artist (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    is_makeup_artist BOOLEAN NOT NULL DEFAULT FALSE,
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

CREATE TABLE mahar (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    mahar_name VARCHAR(255) NOT NULL,
    wedding_ring VARCHAR(50) NOT NULL, -- emas/perak/platinum
    price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

CREATE TABLE other (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID REFERENCES budget(id) ON DELETE CASCADE,
    other_name TEXT NOT NULL, -- souvenirs, invitations, etc.
    other_price INT NOT NULL,
    payment_status VARCHAR(50) NOT NULL CHECK (payment_status IN ('paid', 'unpaid', 'pending')),
    payment_deadline DATE,
    include_setup JSONB DEFAULT '{}' -- JSONB field
);

CREATE TABLE guest (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    address TEXT NOT NULL,
    notes JSONB DEFAULT '{}'
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    google_id VARCHAR(255) UNIQUE NULL, -- Store Google ID (if login via Google)
    provider VARCHAR(50) NOT NULL DEFAULT 'local', -- 'local' or 'google'
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NULL, -- Nullable for Google SSO users
    avatar_url TEXT NULL, -- Profile picture from Google
    role VARCHAR(50) NOT NULL DEFAULT 'user', -- 'user', 'admin', etc.
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    last_login TIMESTAMP NULL
);


CREATE TABLE wedding_organizer (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    contact_wo JSONB DEFAULT '{}'
);



