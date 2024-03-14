package com.tokogue.serviceproduct.config;

import org.apache.solr.client.solrj.SolrClient;
import org.apache.solr.client.solrj.impl.HttpSolrClient;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.concurrent.TimeUnit;

@Configuration
public class SolrConfig {

    @Value("${solr.host}")
    private String solrUrl;

    @Bean
    public SolrClient buildSolrClient() {
        return new HttpSolrClient.Builder(solrUrl)
                .withConnectionTimeout(10000, TimeUnit.MILLISECONDS)
                .withSocketTimeout(60000, TimeUnit.MILLISECONDS)
                .build();
    }

}
