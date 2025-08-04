package com.persico.microservices.fruit_total_pricing.service.model;

import java.math.BigDecimal;

public class FruitTotal {
    private Long id;
    private String fruit;
    private String month;
    private BigDecimal fmp;
    private BigDecimal quantity;
    private BigDecimal total;
    private String port;

    public FruitTotal() {
    }

    public FruitTotal(Long id, String fruit, String month, BigDecimal fmp,
            BigDecimal quantity, BigDecimal total, String port) {
        this.id = id;
        this.fruit = fruit;
        this.month = month;
        this.fmp = fmp;
        this.quantity = quantity;
        this.total = total;
        this.port = port;
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

    public BigDecimal getQuantity() {
        return quantity;
    }

    public void setQuantity(BigDecimal quantity) {
        this.quantity = quantity;
    }

    public BigDecimal getTotal() {
        return total;
    }

    public void setTotal(BigDecimal total) {
        this.total = total;
    }

    public String getPort() {
        return port;
    }

    public void setPort(String port) {
        this.port = port;
    }
}
