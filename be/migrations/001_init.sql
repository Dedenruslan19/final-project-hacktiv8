-- ENUM Definitions
CREATE TYPE user_role AS ENUM ('donor', 'verifikator', 'admin', 'bidder');

CREATE TYPE donation_status AS ENUM (
    'pending',
    'verified_for_auction',
    'verified_for_donation'
);

CREATE TYPE verification_decision AS ENUM ('auction', 'donation');

CREATE TYPE auction_item_status AS ENUM ('scheduled', 'ongoing', 'finished');

CREATE TYPE payment_status AS ENUM ('pending', 'paid', 'failed');