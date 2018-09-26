package io.talha.invoiceapi;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;

@RestController
public class InvoiceController {

    @RequestMapping(value = "/v1/invoice", method = {RequestMethod.GET})
    public String getInvoice() {
        return "Hello Haaris from Spring \n";
    }

    @RequestMapping(value = "/v1/invoice", method = {RequestMethod.POST})
    public String createInvoice() {
        return "Hello Haaris from Spring \n";
    }

}


