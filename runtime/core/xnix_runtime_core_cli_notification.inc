static void
print_notification_application(const XnixRuntimeNotificationPlan *plan)
{
  fputs("{", stdout);
  print_string_field("id", plan->application->id);
  fputs(",", stdout);
  print_string_field("name", plan->application->name);
  fputs(",", stdout);
  print_string_field("icon", plan->application->icon);
  fputs(",", stdout);
  print_string_field("runtime_mode", plan->application->runtime_mode);
  fputs(",", stdout);
  print_string_field("desktop_entry", plan->desktop_entry);
  fputs("}", stdout);
}

static void
print_notification_record(const XnixRuntimeNotificationPlan *plan)
{
  fputs("{", stdout);
  print_string_field("id", plan->notification_id);
  fputs(",", stdout);
  print_string_field("title", plan->title);
  fputs(",", stdout);
  print_string_field("body", plan->body);
  fputs(",", stdout);
  print_string_field("urgency", plan->urgency);
  fputs(",", stdout);
  print_string_field("category", plan->category);
  fputs(",", stdout);
  print_string_field("desktop_entry", plan->desktop_entry);
  fputs(",", stdout);
  fputs("\"actions\":", stdout);
  print_string_array(plan->actions, plan->action_count);
  fputs(",", stdout);
  fputs("\"action_count\":", stdout);
  printf("%zu", plan->action_count);
  fputs("}", stdout);
}

static void
print_notification_safety(const XnixRuntimeNotificationPlan *plan)
{
  fputs("{", stdout);
  fputs("\"user_visible\":", stdout);
  print_bool(plan->user_visible);
  fputs(",", stdout);
  fputs("\"requires_user_review\":", stdout);
  print_bool(plan->requires_user_review);
  fputs(",", stdout);
  fputs("\"action_execution_enabled\":", stdout);
  print_bool(plan->action_execution_enabled);
  fputs(",", stdout);
  fputs("\"repair_execution_enabled\":", stdout);
  print_bool(plan->repair_execution_enabled);
  fputs(",", stdout);
  fputs("\"settings_persistence_enabled\":", stdout);
  print_bool(plan->settings_persistence_enabled);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(plan->host_root_modified);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(plan->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_notification_plan_record(const XnixRuntimeNotificationPlan *plan)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("plan_type", "notification-plan");
  fputs(",", stdout);
  print_string_field("desktop", "KDE Plasma");
  fputs(",", stdout);
  print_string_field("event_type", plan->event_type);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(plan->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(plan->kde_policy_owner);
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_notification_application(plan);
  fputs(",", stdout);
  fputs("\"notification\":", stdout);
  print_notification_record(plan);
  fputs(",", stdout);
  fputs("\"safety\":", stdout);
  print_notification_safety(plan);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", plan->summary);
  fputs("}", stdout);
}

static int
print_notification_plan_for_application(const char *application_id, const char *event_type)
{
  XnixRuntimeNotificationPlan plan;

  if (!xnix_runtime_notification_plan(application_id, event_type, &plan)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or notification plan catalog\n");
    return 64;
  }

  print_notification_plan_record(&plan);
  fputs("\n", stdout);
  return 0;
}
