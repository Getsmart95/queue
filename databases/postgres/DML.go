package postgres

const RolesSeeder  =  `insert into roles(name, display_name, description) values('admin', 'ADMIN', 'Администратор'),
																				('manager', 'MANAGER', 'Управляющий'),
																				('user', 'USER', 'Пользователь')
																			on conflict (name) do nothing`
const AddAdmin =  `insert into users(name, surname, login, password, email, phone, status) values('admin', 'admin', 'admin', '21232f297a57a5a743894a0e4a801fc3', 'admin@admin.tj', '123456789', true)
					on conflict (login) do nothing`
const AddUser  =  `insert into users(name, surname, login, password, email, phone, status) values($1, $2, $3, $4, $5, $6, $7)`
const AddUserRole  =  `insert into userRole(role_id, user_id) values($1, $2)`
const GetUserByLogin  =  `select * from users where login = ($1)`
const GetUserByEmail  =  `select * from users where email = ($1)`
const GetAllRoles  =  `select * from roles`
const AddCity  =  `insert into cities(name) values($1)`
const GetAllCities  =  `select * from cities`
const AddBranch  =  `insert into branches(address, city_id) values($1, $2)`
const GetBranchByCity  =  `select * from branches where city_id = ($1)`
const AddPurpose  =  `insert into purposes(name) values($1)`
const GetPurposes  =  `select * from purposes`
const AddTime  =  `insert into times(name) values($1)`
const GetTimes  =  `select * from times`