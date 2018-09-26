package invoice;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestParam;
import entity.Invoice;

@RestController
public class InvoiceController {

	private String api_version = "v1";
	private Invoice default_invoice = new Invoice(0, 0, 0, 0);

	@RequestMapping("/")
	public String index() {
		return "Greetings from Invoice Api";
	}

	// Get Controller for Invoice
	@RequestMapping(method=RequestMethod.GET, value="/v1/invoice/{id}")
    public Invoice getInvoice(@PathVariable("id") long id) {
        return this.default_invoice;
    }

    // Post Controller for Invoice
	@RequestMapping(method=RequestMethod.POST, value="/v1/invoice")
    public Invoice createInvoice(
    	@RequestParam(value="id") long id,
    	@RequestParam(value="invoice_id") long invoice_id, 
    	@RequestParam(value="po_number") long po_number,
    	@RequestParam(value="amount_cents") long amount_cents
		) {
        return new Invoice(id, 
        	invoice_id, po_number, amount_cents);
    }
}