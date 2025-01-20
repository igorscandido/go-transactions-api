CREATE TABLE payment (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    paymentBaseCurrency VARCHAR(3) NOT NULL,
    paymentBaseAmount NUMERIC(20, 8) NOT NULL,
    gateway VARCHAR(255) NOT NULL,
    gatewayTransactionatedCurrency VARCHAR(3) NOT NULL,
    gatewayTransactionatedAmount NUMERIC(20, 8) NOT NULL,
    currencyConversionRate NUMERIC(20, 8) NOT NULL,
    creationRequestTime TIMESTAMPTZ NOT NULL
);
