import { createClient } from "@supabase/supabase-js";

const SUPABASE_URL = "https://agxslvqovpnwfpocobuw.supabase.co";
const SUPABASE_KEY =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImFneHNsdnFvdnBud2Zwb2NvYnV3Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTk2NTM3NTEsImV4cCI6MjAzNTIyOTc1MX0.DxGvDrHapjN1ir-yMPTEBt8IRwxvvCjNnZ8MQojduko";

const supabase = createClient(SUPABASE_URL, SUPABASE_KEY);
supabase.auth.signUp();
