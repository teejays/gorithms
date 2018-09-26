package repository;

import java.util.List;
import entity.Invoice;

import org.springframework.data.repository.CrudRepository;

public interface InvoiceRepository extends CrudRepository<Invoice, Long> {

    List<Invoice> findByInvoiceId(long invoice_id);
}