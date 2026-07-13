static void
print_compatibility_action(const XnixRuntimeCompatibilityAction *action)
{
  fputs("{", stdout);
  print_string_field("id", action->id);
  fputs(",", stdout);
  print_string_field("source_type", action->source_type);
  fputs(",", stdout);
  print_string_field("status", action->status);
  fputs(",", stdout);
  print_string_field("priority", action->priority);
  fputs(",", stdout);
  print_string_field("title", action->title);
  fputs(",", stdout);
  print_string_field("summary", action->summary);
  fputs(",", stdout);
  fputs("\"user_review_required\":", stdout);
  print_bool(action->user_review_required);
  fputs(",", stdout);
  print_string_field("runtime_gate", action->runtime_gate);
  fputs(",", stdout);
  print_string_field("next_step", action->next_step);
  fputs(",", stdout);
  fputs("\"execution_enabled\":", stdout);
  print_bool(action->execution_enabled);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(action->backend_details_exposed);
  fputs("}", stdout);
}

static void
print_compatibility_actions(const XnixRuntimeCompatibilityActionQueue *queue)
{
  fputc('[', stdout);
  for (size_t index = 0; index < queue->action_count; index++) {
    if (index > 0) {
      fputs(",", stdout);
    }
    print_compatibility_action(&queue->actions[index]);
  }
  fputc(']', stdout);
}

static void
print_compatibility_action_queue_record(
  const XnixRuntimeCompatibilityActionQueue *queue
)
{
  fputs("{", stdout);
  printf("\"version\":\"%s\",", xnix_runtime_version());
  print_string_field("queue_type", "compatibility-center-action-queue");
  fputs(",", stdout);
  fputs("\"application\":", stdout);
  print_install_application(queue->application);
  fputs(",", stdout);
  fputs("\"runtime_owned\":", stdout);
  print_bool(queue->runtime_owned);
  fputs(",", stdout);
  fputs("\"kde_policy_owner\":", stdout);
  print_bool(queue->kde_policy_owner);
  fputs(",", stdout);
  print_string_field("surface", "Compatibility Center");
  fputs(",", stdout);
  fputs("\"action_count\":", stdout);
  printf("%zu,", queue->action_count);
  fputs("\"pending_action_count\":", stdout);
  printf("%zu,", queue->pending_action_count);
  fputs("\"user_review_required_count\":", stdout);
  printf("%zu,", queue->user_review_required_count);
  fputs("\"execution_enabled\":", stdout);
  print_bool(queue->execution_enabled);
  fputs(",", stdout);
  fputs("\"repair_execution_enabled\":", stdout);
  print_bool(queue->repair_execution_enabled);
  fputs(",", stdout);
  fputs("\"settings_persistence_enabled\":", stdout);
  print_bool(queue->settings_persistence_enabled);
  fputs(",", stdout);
  fputs("\"host_root_modified\":", stdout);
  print_bool(queue->host_root_modified);
  fputs(",", stdout);
  fputs("\"network_required\":", stdout);
  print_bool(queue->network_required);
  fputs(",", stdout);
  fputs("\"backend_details_exposed\":", stdout);
  print_bool(queue->backend_details_exposed);
  fputs(",", stdout);
  fputs("\"actions\":", stdout);
  print_compatibility_actions(queue);
  fputs(",", stdout);
  fputs("\"blocked_actions\":", stdout);
  print_string_array(queue->blocked_actions, queue->blocked_action_count);
  fputs(",", stdout);
  print_string_field("desktop_safe_summary", queue->summary);
  fputs("}", stdout);
}

static int
print_compatibility_action_queue_for_application(const char *application_id)
{
  XnixRuntimeCompatibilityActionQueue queue;

  if (!xnix_runtime_compatibility_action_queue(application_id, &queue)) {
    fprintf(stderr, "xnix-runtime-core: unknown application or action queue catalog\n");
    return 64;
  }

  print_compatibility_action_queue_record(&queue);
  fputs("\n", stdout);
  return 0;
}
