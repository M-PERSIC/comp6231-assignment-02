package com.persico.microservices.fruit_month_pricing.service;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.persico.microservices.fruit_month_pricing.service.service.FruitPriceService;

@SpringBootApplication
public class FruitMonthPriceApplication implements CommandLineRunner {
    
    private Logger logger = LoggerFactory.getLogger(FruitMonthPriceApplication.class);
    
    @Autowired
    private FruitPriceService fruitPriceService;
    
    public static void main(String[] args) {
        SpringApplication.run(FruitMonthPriceApplication.class, args);
    }
    
    @Override
    public void run(String... args) throws Exception {
        logger.info("Starting Fruit Month Price Service");
        long totalRecords = fruitPriceService.getTotalRecords();
        logger.info("Fruit Month Price Service started successfully with {} records", totalRecords);
    }
}
