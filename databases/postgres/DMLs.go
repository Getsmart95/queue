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