package entity;

import java.time.LocalDateTime;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;


@Entity(name = "tb_invoice")
public class Invoice {

    @Id
    @GeneratedValue(strategy=GenerationType.AUTO)
    private long id;
    private long invoice_number, po_number, amount_cents;
    private LocalDateTime date_created;

    public Invoice(long id, long invoice_number, long po_number, long amount_cents) {
        this.id = id;
        this.invoice_number = invoice_number;
        this.po_number = po_number;
        this.amount_cents = amount_cents;

        this.date_created = LocalDateTime.now();
    }

    @Override
    public String toString() {
        return String.format(
            "Invoice[id=%d, invoice_number=%d, po_number=%d, amount_cents=%d]",
             this.id, this.invoice_number, this.po_number, this.amount_cents);
    }

}

