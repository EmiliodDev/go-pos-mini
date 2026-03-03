-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD CONSTRAINT fk_copanies FOREIGN KEY (company_id) REFERENCES companies(company_id) DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE users ADD CONSTRAINT fk_branches FOREIGN KEY (branch_id) REFERENCES branches(branch_id) DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE branches ADD CONSTRAINT fk_companies FOREIGN KEY (company_id) REFERENCES companies(company_id) DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE branches ADD CONSTRAINT fk_warehouses FOREIGN KEY (warehouse_id) REFERENCES warehouses(warehouse_id) DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE inventories ADD CONSTRAINT fk_warehouses FOREIGN KEY (warehouse_id) REFERENCES warehouses(warehouse_id) DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE inventories ADD CONSTRAINT fk_products FOREIGN KEY (product_id) REFERENCES products(product_id) DEFERRABLE INITIALLY IMMEDIATE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_companies;
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_branches;
ALTER TABLE branches DROP CONSTRAINT IF EXISTS fk_companies;
ALTER TABLE branches DROP CONSTRAINT IF EXISTS fk_warehouses;
ALTER TABLE inventories DROP CONSTRAINT IF EXISTS fk_warehouses;
ALTER TABLE inventories DROP CONSTRAINT IF EXISTS fk_products;
-- +goose StatementEnd
