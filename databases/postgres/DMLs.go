package postgres

const SetTimeZone  =  `SET timezone='Asia/Dushanbe';`
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
const AddQueue  =  `insert into queues(queue_code, user_id, city_id, branch_id, purpose_id, time_id, status, date)
					values($1, $2, $3, $4, $5, $6, $7, $8)`
const GetQueuesByDate  =  `select * from queues where date = ($1)`
const GetQueuesByTime  =  `select * from queues where time_id = ($1)`
const GetQueuesByStatus  =  `select * from queues where status = ($1)`
const GetQueuesByUser  =  `select * from queues where user_id = ($1)`
const UpdateQueue  =  `update queues set queue_code = ($1), user_id = ($2), city_id = ($3), branch_id = ($4), purpose_id = ($5), time_id = ($6), status = ($7), date = ($8) where id = ($9)`
const QueueChangeStatus  =  `update queues set status = ($1) where id = ($2)`
const UpdateUser  =  `update users set name = ($1), surname = ($2), email = ($3), phone = ($4), status = ($5) where id = ($6)`
const GetRoleByUser  =  `select ur.role_id, u.id, r.name 
						 from userrole ur, users u, roles r 
						 where u.id = ur.user_id and r.id = ur.role_id and u.login = ($1)`