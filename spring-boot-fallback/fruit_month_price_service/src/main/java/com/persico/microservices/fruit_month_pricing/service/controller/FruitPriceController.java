package com.persico.microservices.fruit_month_pricing.service.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import com.persico.microservices.fruit_month_pricing.service.model.FruitPrice;
import com.persico.microservices.fruit_month_pricing.service.repository.FruitPriceRepository;

@RestController
public class FruitPriceController {
    
    private Logger logger = LoggerFactory.getLogger(FruitPriceController.class);
    
    @Autowired
    private FruitPriceRepository repository;
    
    @Autowired
    private Environment environment;
    
    @GetMapping("/fruit-price/fruit/{fruit}/month/{month}")
    public FruitPrice retrieveFruitPrice(
            @PathVariable String fruit,
            @PathVariable String month) {
        logger.info("retrieveFruitPrice called with {} to {}", fruit, month);
        String fruitKey = capitalizeFirst(fruit.toLowerCase());
        String monthKey = capitalizeFirst(month.toLowerCase());
        FruitPrice fruitPrice = repository.findByFruitAndMonth(fruitKey, monthKey);
        if (fruitPrice == null) {
            throw new RuntimeException("Unable to find data for " + fruitKey + " in " + monthKey);
        }
        String port = environment.getProperty("local.server.port");
        fruitPrice.setEnvironment(port);
        return fruitPrice;
    }
    
    private String capitalizeFirst(String str) {
        if (str == null || str.isEmpty()) {
            return str;
        }
        return str.substring(0, 1).toUpperCase() + str.substring(1);
    }
}
