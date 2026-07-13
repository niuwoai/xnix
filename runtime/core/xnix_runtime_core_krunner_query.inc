static void
normalize_krunner_query(const char *input, char *output, size_t output_size)
{
  if (output_size == 0) {
    return;
  }

  size_t written = 0;
  bool previous_space = true;
  for (const unsigned char *cursor = (const unsigned char *)input;
       cursor != NULL && *cursor != '\0' && written + 1 < output_size;
       cursor++) {
    if (isspace(*cursor)) {
      if (!previous_space) {
        output[written++] = ' ';
        previous_space = true;
      }
      continue;
    }

    output[written++] = (char)tolower(*cursor);
    previous_space = false;
  }

  if (written > 0 && output[written - 1] == ' ') {
    written--;
  }
  output[written] = '\0';
}

static bool
krunner_query_matches_sample(const char *normalized_query)
{
  if (normalized_query == NULL || normalized_query[0] == '\0') {
    return false;
  }

  return strcmp(normalized_query, "notepad") == 0 ||
         strcmp(normalized_query, "sample notepad") == 0 ||
         strcmp(normalized_query, "org.xnix.sample.notepad") == 0 ||
         strcmp(normalized_query, "open notepad") == 0 ||
         strcmp(normalized_query, "launch notepad") == 0 ||
         strcmp(normalized_query, "start notepad") == 0 ||
         strcmp(normalized_query, "run notepad") == 0 ||
         strcmp(normalized_query, "txt") == 0 ||
         strcmp(normalized_query, ".txt") == 0 ||
         strcmp(normalized_query, "open txt") == 0 ||
         strcmp(normalized_query, "launch txt") == 0 ||
         strcmp(normalized_query, "open text") == 0 ||
         strstr(normalized_query, "notepad") != NULL;
}

static size_t
krunner_relevance_for_query(const char *normalized_query)
{
  if (strcmp(normalized_query, "sample notepad") == 0 ||
      strcmp(normalized_query, "notepad") == 0) {
    return 100;
  }
  if (strcmp(normalized_query, "open notepad") == 0 ||
      strcmp(normalized_query, "launch notepad") == 0 ||
      strcmp(normalized_query, "start notepad") == 0 ||
      strcmp(normalized_query, "run notepad") == 0) {
    return 90;
  }
  if (strcmp(normalized_query, "txt") == 0 ||
      strcmp(normalized_query, ".txt") == 0 ||
      strcmp(normalized_query, "open txt") == 0 ||
      strcmp(normalized_query, "launch txt") == 0) {
    return 85;
  }

  return 75;
}

bool
xnix_runtime_krunner_query_plan(
  const char *query,
  XnixRuntimeKRunnerQueryPlan *plan
)
{
  if (plan == NULL) {
    return false;
  }

  char normalized_query[160];
  normalize_krunner_query(query == NULL ? "" : query, normalized_query, sizeof(normalized_query));

  plan->query_type = "krunner-query-plan";
  plan->entry_point = "krunner";
  plan->desktop = "KDE Plasma";
  plan->query = query == NULL ? "" : query;
  plan->match_count = 0;
  plan->runtime_owned = true;
  plan->kde_policy_owner = false;
  plan->runtime_owned_launch = true;
  plan->query_execution_enabled = false;
  plan->backend_launch_enabled = false;
  plan->host_root_modified = false;
  plan->backend_details_exposed = false;
  plan->summary = "KRunner query planning is Runtime-owned and returns safe launcher actions only.";

  if (!krunner_query_matches_sample(normalized_query)) {
    return true;
  }

  const XnixRuntimeApplication *application =
    xnix_runtime_find_application("org.xnix.sample.notepad");
  if (application == NULL) {
    return true;
  }

  XnixRuntimeKRunnerMatch *match = &plan->matches[0];
  match->application = application;
  match->runner_id = "xnix.compatibility.org.xnix.sample.notepad";
  match->subtitle = "Open as a normal Linux application";
  match->mode_label = "Automatic";
  match->action_type = "runtime-launch";
  match->desktop_entry_id = "org.xnix.sample.notepad.desktop";
  match->action_argv[0] = "xnix-compat-launch";
  match->action_argv[1] = "--app";
  match->action_argv[2] = application->id;
  match->action_argc = 3;
  match->supported_extensions[0] = application->primary_extension;
  match->supported_extensions[1] = ".log";
  match->supported_extension_count = 2;
  match->relevance_percent = krunner_relevance_for_query(normalized_query);
  match->runtime_owned_launch = true;
  match->backend_details_exposed = false;
  plan->match_count = 1;

  return true;
}
