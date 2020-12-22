package postgres

const RolesSeeder  =  `insert into roles(name, display_name, description) values('admin', 'ADMIN', 'Администратор'),
																				('manager', 'MANAGER', 'Управляющий'),
																				('user', 'USER', 'Пользователь')`
const AddAdmin =  `insert into users(name, surname, login, password, email, phone, status) values('admin', 'admin', 'admin', 'admin', 'admin@admin.tj', '123456789', 'admin')`
const AddUser  =  `insert into users(name, surname, login, password, email, phone, status) values($1, $2, $3, $4, $5, $6, $7)`
const AddUserRole  =  `insert into userRole(role_id, user_id) values($1, $2)`
const GetUserByLogin  =  `select * from users where login = ($1)`
const GetAllRoles  =  `select * from roles`