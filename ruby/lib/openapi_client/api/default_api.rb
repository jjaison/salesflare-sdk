=begin
#app.salesflare.com

#No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

The version of the OpenAPI document: 1.0.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 7.3.0

=end

require 'cgi'

module OpenapiClient
  class DefaultApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # /messenger/web/metrics
    # **Host**: http://api-iam.intercom.io
    # @param app_id [String] 
    # @param v [String] 
    # @param g [String] 
    # @param s [String] 
    # @param r [String] 
    # @param platform [String] 
    # @param installation_type [String] 
    # @param idempotency_key [String] 
    # @param internal [String] 
    # @param is_intersection_booted [String] 
    # @param page_title [String] 
    # @param user_active_company_id [String] 
    # @param metrics [String] 
    # @param logs [String] 
    # @param op_metrics [String] 
    # @param hc_metrics [String] 
    # @param referer [String] 
    # @param anonymous_session [String] 
    # @param device_identifier [String] 
    # @param [Hash] opts the optional parameters
    # @return [nil]
    def messenger_web_metrics_post(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier, opts = {})
      messenger_web_metrics_post_with_http_info(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier, opts)
      nil
    end

    # /messenger/web/metrics
    # **Host**: http://api-iam.intercom.io
    # @param app_id [String] 
    # @param v [String] 
    # @param g [String] 
    # @param s [String] 
    # @param r [String] 
    # @param platform [String] 
    # @param installation_type [String] 
    # @param idempotency_key [String] 
    # @param internal [String] 
    # @param is_intersection_booted [String] 
    # @param page_title [String] 
    # @param user_active_company_id [String] 
    # @param metrics [String] 
    # @param logs [String] 
    # @param op_metrics [String] 
    # @param hc_metrics [String] 
    # @param referer [String] 
    # @param anonymous_session [String] 
    # @param device_identifier [String] 
    # @param [Hash] opts the optional parameters
    # @return [Array<(nil, Integer, Hash)>] nil, response status code and response headers
    def messenger_web_metrics_post_with_http_info(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DefaultApi.messenger_web_metrics_post ...'
      end
      # verify the required parameter 'app_id' is set
      if @api_client.config.client_side_validation && app_id.nil?
        fail ArgumentError, "Missing the required parameter 'app_id' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'v' is set
      if @api_client.config.client_side_validation && v.nil?
        fail ArgumentError, "Missing the required parameter 'v' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'g' is set
      if @api_client.config.client_side_validation && g.nil?
        fail ArgumentError, "Missing the required parameter 'g' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 's' is set
      if @api_client.config.client_side_validation && s.nil?
        fail ArgumentError, "Missing the required parameter 's' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'r' is set
      if @api_client.config.client_side_validation && r.nil?
        fail ArgumentError, "Missing the required parameter 'r' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'platform' is set
      if @api_client.config.client_side_validation && platform.nil?
        fail ArgumentError, "Missing the required parameter 'platform' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'installation_type' is set
      if @api_client.config.client_side_validation && installation_type.nil?
        fail ArgumentError, "Missing the required parameter 'installation_type' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'idempotency_key' is set
      if @api_client.config.client_side_validation && idempotency_key.nil?
        fail ArgumentError, "Missing the required parameter 'idempotency_key' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'internal' is set
      if @api_client.config.client_side_validation && internal.nil?
        fail ArgumentError, "Missing the required parameter 'internal' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'is_intersection_booted' is set
      if @api_client.config.client_side_validation && is_intersection_booted.nil?
        fail ArgumentError, "Missing the required parameter 'is_intersection_booted' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'page_title' is set
      if @api_client.config.client_side_validation && page_title.nil?
        fail ArgumentError, "Missing the required parameter 'page_title' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'user_active_company_id' is set
      if @api_client.config.client_side_validation && user_active_company_id.nil?
        fail ArgumentError, "Missing the required parameter 'user_active_company_id' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'metrics' is set
      if @api_client.config.client_side_validation && metrics.nil?
        fail ArgumentError, "Missing the required parameter 'metrics' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'logs' is set
      if @api_client.config.client_side_validation && logs.nil?
        fail ArgumentError, "Missing the required parameter 'logs' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'op_metrics' is set
      if @api_client.config.client_side_validation && op_metrics.nil?
        fail ArgumentError, "Missing the required parameter 'op_metrics' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'hc_metrics' is set
      if @api_client.config.client_side_validation && hc_metrics.nil?
        fail ArgumentError, "Missing the required parameter 'hc_metrics' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'referer' is set
      if @api_client.config.client_side_validation && referer.nil?
        fail ArgumentError, "Missing the required parameter 'referer' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'anonymous_session' is set
      if @api_client.config.client_side_validation && anonymous_session.nil?
        fail ArgumentError, "Missing the required parameter 'anonymous_session' when calling DefaultApi.messenger_web_metrics_post"
      end
      # verify the required parameter 'device_identifier' is set
      if @api_client.config.client_side_validation && device_identifier.nil?
        fail ArgumentError, "Missing the required parameter 'device_identifier' when calling DefaultApi.messenger_web_metrics_post"
      end
      # resource path
      local_var_path = '/messenger/web/metrics'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['text/html'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/x-www-form-urlencoded'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}
      form_params['app_id'] = app_id
      form_params['v'] = v
      form_params['g'] = g
      form_params['s'] = s
      form_params['r'] = r
      form_params['platform'] = platform
      form_params['installation_type'] = installation_type
      form_params['Idempotency-Key'] = idempotency_key
      form_params['internal'] = internal
      form_params['is_intersection_booted'] = is_intersection_booted
      form_params['page_title'] = page_title
      form_params['user_active_company_id'] = user_active_company_id
      form_params['metrics'] = metrics
      form_params['logs'] = logs
      form_params['op_metrics'] = op_metrics
      form_params['hc_metrics'] = hc_metrics
      form_params['referer'] = referer
      form_params['anonymous_session'] = anonymous_session
      form_params['device_identifier'] = device_identifier

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type]

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"DefaultApi.messenger_web_metrics_post",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DefaultApi#messenger_web_metrics_post\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # /tasks
    # **Host**: http://api.salesflare.com
    # @param [Hash] opts the optional parameters
    # @option opts [String] :limit 
    # @option opts [String] :q 
    # @return [Array<TasksGet200ResponseInner>]
    def tasks_get(opts = {})
      data, _status_code, _headers = tasks_get_with_http_info(opts)
      data
    end

    # /tasks
    # **Host**: http://api.salesflare.com
    # @param [Hash] opts the optional parameters
    # @option opts [String] :limit 
    # @option opts [String] :q 
    # @return [Array<(Array<TasksGet200ResponseInner>, Integer, Hash)>] Array<TasksGet200ResponseInner> data, response status code and response headers
    def tasks_get_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DefaultApi.tasks_get ...'
      end
      # resource path
      local_var_path = '/tasks'

      # query parameters
      query_params = opts[:query_params] || {}
      query_params[:'limit'] = opts[:'limit'] if !opts[:'limit'].nil?
      query_params[:'q'] = opts[:'q'] if !opts[:'q'].nil?

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'Array<TasksGet200ResponseInner>'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apikey cookie  hjsessionuser 374224', 'apikey cookie  gid', 'apikey cookie  hjsession 374224', 'apikey cookie ajs user id', 'apikey cookie intercom-id-nhqftro2', 'apikey cookie intercom-session-nhqftro2', 'apikey cookie intercom-device-id-nhqftro2', 'apikey cookie salesflare-session', 'apikey cookie ajs anonymous id', 'apikey header cookie']

      new_options = opts.merge(
        :operation => :"DefaultApi.tasks_get",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:GET, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DefaultApi#tasks_get\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # /tasks
    # **Host**: http://api.salesflare.com
    # @param [Hash] opts the optional parameters
    # @return [TasksPost200Response]
    def tasks_post(opts = {})
      data, _status_code, _headers = tasks_post_with_http_info(opts)
      data
    end

    # /tasks
    # **Host**: http://api.salesflare.com
    # @param [Hash] opts the optional parameters
    # @return [Array<(TasksPost200Response, Integer, Hash)>] TasksPost200Response data, response status code and response headers
    def tasks_post_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DefaultApi.tasks_post ...'
      end
      # resource path
      local_var_path = '/tasks'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json;charset=UTF-8'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'TasksPost200Response'

      # auth_names
      auth_names = opts[:debug_auth_names] || ['apikey cookie  hjsessionuser 374224', 'apikey cookie  gid', 'apikey cookie  hjsession 374224', 'apikey cookie ajs user id', 'apikey cookie intercom-id-nhqftro2', 'apikey cookie intercom-session-nhqftro2', 'apikey cookie intercom-device-id-nhqftro2', 'apikey cookie salesflare-session', 'apikey cookie ajs anonymous id', 'apikey header cookie']

      new_options = opts.merge(
        :operation => :"DefaultApi.tasks_post",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DefaultApi#tasks_post\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # /v1/m
    # **Host**: http://api.segment.io
    # @param [Hash] opts the optional parameters
    # @return [V1MPost200Response]
    def v1_m_post(opts = {})
      data, _status_code, _headers = v1_m_post_with_http_info(opts)
      data
    end

    # /v1/m
    # **Host**: http://api.segment.io
    # @param [Hash] opts the optional parameters
    # @return [Array<(V1MPost200Response, Integer, Hash)>] V1MPost200Response data, response status code and response headers
    def v1_m_post_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: DefaultApi.v1_m_post ...'
      end
      # resource path
      local_var_path = '/v1/m'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['text/plain'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'V1MPost200Response'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"DefaultApi.v1_m_post",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: DefaultApi#v1_m_post\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
