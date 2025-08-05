package com.persico.microservices.fruit_month_pricing.service.model;

import java.math.BigDecimal;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;

@Entity
public class FruitPrice {
    
    @Id
    private Long id;
    
    @Column(name = "fruit_name")
    private String fruit;
    
    @Column(name = "month_name")
    private String month;
    
    private BigDecimal fmp;
    
    private String environment;
    
    public FruitPrice() {}
    
    public FruitPrice(Long id, String fruit, String month, BigDecimal fmp, String environment) {
        super();
        this.id = id;
        this.fruit = fruit;
        this.month = month;
        this.fmp = fmp;
        this.environment = environment;
    }
    
    public Long getId() {
        return id;
    }
    
    public void setId(Long id) {
        this.id = id;
    }
    
    public String getFruit() {
        return fruit;
    }
    
    public void setFruit(String fruit) {
        this.fruit = fruit;
    }
    
    public String getMonth() {
        return month;
    }
    
    public void setMonth(String month) {
        this.month = month;
    }
    
    public BigDecimal getFmp() {
        return fmp;
    }
    
    public void setFmp(BigDecimal fmp) {
        this.fmp = fmp;
    }
    
    public String getEnvironment() {
        return environment;
    }
    
    public void setEnvironment(String environment) {
        this.environment = environment;
    }
}
