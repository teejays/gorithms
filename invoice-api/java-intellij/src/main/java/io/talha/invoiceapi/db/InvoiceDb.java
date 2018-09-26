package io.talha.invoiceapi.db;

import io.talha.invoiceapi.entity.Invoice;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

public interface InvoiceDb extends CrudRepository<Invoice, Long> { }
