package com.persico.microservices.fruit_total_pricing.service.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

import com.persico.microservices.fruit_total_pricing.service.model.FruitTotal;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.Map;

@RestController
public class FruitTotalPriceController {
    @Autowired
    private Environment environment;
    @Autowired
    private RestTemplate restTemplate;
    @GetMapping("/fruit-total/fruit/{fruitName}/month/{monthName}/quantity/{quantity}")
    public FruitTotal calculateFruitTotal(
            @PathVariable String fruitName,
            @PathVariable String monthName,
            @PathVariable BigDecimal quantity) {
        String fmpPort = environment.getProperty("fmp.port", "8000");
        String fmpServiceUrl = "http://localhost:" + fmpPort + "/fruit-price/fruit/{fruit}/month/{month}";
        Map<String, String> uriVariables = new HashMap<>();
        uriVariables.put("fruit", fruitName);
        uriVariables.put("month", monthName);
        FruitTotal fruitMonthPrice = restTemplate.getForObject(
                fmpServiceUrl,
                FruitTotal.class,
                uriVariables);
        if (fruitMonthPrice == null) {
            throw new RuntimeException("Unable to find data for " + fruitName + " in " + monthName);
        }
        String port = environment.getProperty("local.server.port");
        BigDecimal totalPrice = quantity.multiply(fruitMonthPrice.getFmp());
        return new FruitTotal(
                fruitMonthPrice.getId(),
                fruitName,
                monthName,
                fruitMonthPrice.getFmp(),
                quantity,
                totalPrice,
                port);
    }
}
