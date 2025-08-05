package com.persico.microservices.fruit_month_pricing.service.service;

import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.persico.microservices.fruit_month_pricing.service.model.FruitPrice;
import com.persico.microservices.fruit_month_pricing.service.repository.FruitPriceRepository;

@Service
public class FruitPriceService {
    
    private Logger logger = LoggerFactory.getLogger(FruitPriceService.class);
    
    @Autowired
    private FruitPriceRepository repository;
    
    public FruitPrice getFruitPrice(String fruit, String month) {
        logger.info("Getting fruit price for {} in {}", fruit, month);
        return repository.findByFruitAndMonth(fruit, month);
    }
    
    public List<FruitPrice> getAllPricesForFruit(String fruit) {
        logger.info("Getting all prices for fruit: {}", fruit);
        return repository.findByFruit(fruit);
    }
    
    public List<FruitPrice> getAllPricesForMonth(String month) {
        logger.info("Getting all prices for month: {}", month);
        return repository.findByMonth(month);
    }
    
    public long getTotalRecords() {
        return repository.count();
    }
}
