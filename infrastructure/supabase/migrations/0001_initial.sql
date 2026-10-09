-- Optional cloud control-plane foundation. The local-first app does not require Supabase.
create extension if not exists pgcrypto;
create table if not exists public.profiles (id uuid primary key references auth.users(id) on delete cascade, display_name text, created_at timestamptz not null default now());
create table if not exists public.projects (id uuid primary key default gen_random_uuid(),owner_id uuid not null references auth.users(id),name text not null,created_at timestamptz not null default now());
create table if not exists public.goals (id uuid primary key default gen_random_uuid(),project_id uuid not null references public.projects(id) on delete cascade,objective text not null,criteria text,status text not null default 'planned',created_at timestamptz not null default now());
alter table public.profiles enable row level security;
alter table public.projects enable row level security;
alter table public.goals enable row level security;
create policy "profile owner" on public.profiles for all to authenticated using (id=auth.uid()) with check(id=auth.uid());
create policy "project owner" on public.projects for all to authenticated using (owner_id=auth.uid()) with check(owner_id=auth.uid());
create policy "goal project owner" on public.goals for all to authenticated using(exists(select 1 from public.projects p where p.id=project_id and p.owner_id=auth.uid())) with check(exists(select 1 from public.projects p where p.id=project_id and p.owner_id=auth.uid()));
