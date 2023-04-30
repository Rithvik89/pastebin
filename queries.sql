-- name: CreatePaste :exec
insert into pastes (
  shortlink,
  expiration_length_in_minutes,
  created_at,
  paste_path
) Values (?,?,?,?);

-- name: GetPaste :one
select paste_path,expiration_length_in_minutes,created_at from pastes
where shortlink = (?);