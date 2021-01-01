package postgres


const RolesSeeder  =  `insert into roles(name, display_name, description) values('admin', 'ADMIN', 'Администратор'),
																				('manager', 'MANAGER', 'Управляющий'),
																				('user', 'USER', 'Пользователь')
																			on conflict (name) do nothing`
const UserSeeder =  `insert into users(name, surname, login, password, email, phone, status) values('admin', 'admin', 'admin', '$2a$10$BnsDRa0/xI3y.OMer1ANJend3zxe6hEmPhAcx6xXri2h6bAwO6dme', 'admin@admin.tj', '123456789', true)
					on conflict (login) do nothing`
const UserRoleSeeder  =  `insert into userrole(role_id, user_id) values(1, 1)
						on conflict (user_id) do nothing`
const CitiesSeeder  =  `insert into cities(name) values('Душанбе'),('Бохтар'),('Худжант'),('Кулоб'),('Вахдат')
					on conflict (name) do nothing`
const BranchesSeeder  =  `insert into branches(address, city_id) values('Ул.Борбади', 1)
					on conflict (address) do nothing`
const PurposesSeeder  =  `insert into purposes(name) values('Заявление'),('Выписка'),('Справка'),('Кредитная карта')
					on conflict (name) do nothing`
const TimesSeeder  =  `insert into times(name) values('9:00'),('9:30'),('10:00'),('10:30'),('11:00'),('11:30	')
					on conflict (name) do nothing`