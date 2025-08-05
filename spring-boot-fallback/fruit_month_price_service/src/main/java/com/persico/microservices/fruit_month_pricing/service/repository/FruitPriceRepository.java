package com.persico.microservices.fruit_month_pricing.service.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import com.persico.microservices.fruit_month_pricing.service.model.FruitPrice;

public interface FruitPriceRepository extends JpaRepository<FruitPrice, Long> {
    
    FruitPrice findByFruitAndMonth(String fruit, String month);
    
    java.util.List<FruitPrice> findByFruit(String fruit);
    java.util.List<FruitPrice> findByMonth(String month);
}
