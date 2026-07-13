static void
print_krunner_action(const XnixRuntimeKRunnerMatch *match)
{
  fputs("{", stdout);
  print_string_field("type", match->action_type);
  fputs(",", stdout);
  print_string_field("desktop_entry_id", match->desktop_entry_id);
  fputs(",", stdout);
  fputs("\"argv\":", stdout);
  print_string_array(match->action_argv, match->action_argc);
  fputs("}", stdout);
}

static void
print_krunner_match(const XnixRuntimeKRunnerMatch *match)
{
  fputs("{", stdout);
  print_string_field("runner_id", match->runner_id);
  fputs(",", stdout);
  print_string_field("application_id", match->application->id);
  fputs(",", stdout);
  print_string_field("name", match->application->name);
  fputs(",", stdout);
  print_string_field("icon", match->application->icon);
  fputs(",", stdout);
  fputs("\"relevance_percent\":", stdout);
  printf("%zu,", match->relevance_percent);
  print_string_field("subtitle", match->subtitle);
  fputs(",", stdout);
  print_string_field("mode_label", match->mode_label);
  fputs(",", stdout);
  fputs("\"supported_extensions\":", stdout);
  print_string_array(match->supported_extensions, match->supported_extension_count);
  fputs(",", stdout);
  fputs("\"runtime_owned_launch\":", stdout);
  print_bool(match->runtime_owned_launch);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(match->backend_details_exposed);
  fputs(",", stdout);
  fputs("\"action\":", stdout);
  print_krunner_action(match);
  fputs("}", stdout);
}

static void
print_krunner_matches(const XnixRuntimeKRunnerQueryPlan *plan)
{
  fputc('[', stdout);
  for (size_t index = 0; index < plan->match_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_krunner_match(&plan->matches[index]);
  }
  fputc(']', stdout);
}

static void
print_krunner_summary(const XnixRuntimeKRunnerQueryPlan *plan)
{
  fputs("{", stdout);
  fputs("\"match_count\":", stdout);
  printf("%zu,", plan->match_count);
  print_string_field("official_desktop", plan->desktop);
  fputs(",", stdout);
  fputs("\"runtime_owned_launch\":", stdout);
  print_bool(plan->runtime_owned_launch);
  fputs(",", stdout);
  fputs("\"query_execution_enabled\":", stdout);
  print_bool(plan->query_execution_enabled);
  fputs(",", stdout);
  fputs("\"backend_launch_enabled\":", stdout);
  print_bool(plan->backend_launch_enabled);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_krunner_query_plan_record(const XnixRuntimeKRunnerQueryPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("query_type", plan->query_type);
  fputs(",", stdout);
  print_string_field("entry_point", plan->entry_point);
  fputs(",", stdout);
  print_string_field("desktop", plan->desktop);
  fputs(",", stdout);
  print_string_field("query", plan->query);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"matches\":", stdout);
  print_krunner_matches(plan);
  fputs(",", stdout);
  fputs("\"summary\":", stdout);
  print_krunner_summary(plan);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_krunner_query_plan_for_query(const char *query)
{
  XnixRuntimeKRunnerQueryPlan plan;

  if (!xnix_runtime_krunner_query_plan(query, &plan)) {
    fprintf(stderr, "xnix-runtime-core: KRunner query plan is unavailable\n");
    return 64;
  }

  print_krunner_query_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
