package io.talha.invoiceapi.entity;


import javax.persistence.Entity;
import java.time.LocalDateTime;

@Entity(name = "tb_invoice")
public class Invoice {
    private long id, invoice_number, po_number, amount_cents;
    private LocalDateTime date_created;
}
