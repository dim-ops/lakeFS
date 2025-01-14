/*
 * lakeFS API
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.lakefs.clients.api.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * ACL
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ACL {
  public static final String SERIALIZED_NAME_PERMISSION = "permission";
  @SerializedName(SERIALIZED_NAME_PERMISSION)
  private String permission;

  public static final String SERIALIZED_NAME_ALL_REPOSITORIES = "all_repositories";
  @SerializedName(SERIALIZED_NAME_ALL_REPOSITORIES)
  private Boolean allRepositories;

  public static final String SERIALIZED_NAME_REPOSITORIES = "repositories";
  @SerializedName(SERIALIZED_NAME_REPOSITORIES)
  private List<String> repositories = null;


  public ACL permission(String permission) {
    
    this.permission = permission;
    return this;
  }

   /**
   * Permission level to give this ACL.  \&quot;Read\&quot;, \&quot;Write\&quot;, \&quot;Super\&quot; and \&quot;Admin\&quot; are all supported. 
   * @return permission
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Permission level to give this ACL.  \"Read\", \"Write\", \"Super\" and \"Admin\" are all supported. ")

  public String getPermission() {
    return permission;
  }


  public void setPermission(String permission) {
    this.permission = permission;
  }


  public ACL allRepositories(Boolean allRepositories) {
    
    this.allRepositories = allRepositories;
    return this;
  }

   /**
   * If true, this ACL applies to all repositories, including those added in future.  Permission \&quot;Admin\&quot; allows changing ACLs, so this is necessarily true for that permission. 
   * @return allRepositories
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "If true, this ACL applies to all repositories, including those added in future.  Permission \"Admin\" allows changing ACLs, so this is necessarily true for that permission. ")

  public Boolean getAllRepositories() {
    return allRepositories;
  }


  public void setAllRepositories(Boolean allRepositories) {
    this.allRepositories = allRepositories;
  }


  public ACL repositories(List<String> repositories) {
    
    this.repositories = repositories;
    return this;
  }

  public ACL addRepositoriesItem(String repositoriesItem) {
    if (this.repositories == null) {
      this.repositories = new ArrayList<String>();
    }
    this.repositories.add(repositoriesItem);
    return this;
  }

   /**
   * Apply this ACL only to these repositories.
   * @return repositories
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Apply this ACL only to these repositories.")

  public List<String> getRepositories() {
    return repositories;
  }


  public void setRepositories(List<String> repositories) {
    this.repositories = repositories;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ACL ACL = (ACL) o;
    return Objects.equals(this.permission, ACL.permission) &&
        Objects.equals(this.allRepositories, ACL.allRepositories) &&
        Objects.equals(this.repositories, ACL.repositories);
  }

  @Override
  public int hashCode() {
    return Objects.hash(permission, allRepositories, repositories);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ACL {\n");
    sb.append("    permission: ").append(toIndentedString(permission)).append("\n");
    sb.append("    allRepositories: ").append(toIndentedString(allRepositories)).append("\n");
    sb.append("    repositories: ").append(toIndentedString(repositories)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

