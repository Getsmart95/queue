package postgres


const RolesSeeder  =  `insert into roles(name, display_name, description) values('admin', 'ADMIN', 'Администратор'),
																				('manager', 'MANAGER', 'Управляющий'),
																				('user', 'USER', 'Пользователь')
																			on conflict (name) do nothing`
const UserSeeder =  `insert into users(name, surname, login, password, email, phone, status) values('admin', 'admin', 'admin', '21232f297a57a5a743894a0e4a801fc3', 'admin@admin.tj', '123456789', true)
					on conflict (login) do nothing`
const CitiesSeeder  =  `insert into cities(name) values('Душанбе'),('Бохтар'),('Худжант'),('Кулоб'),('Вахдат')
					on conflict (name) do nothing`
const BranchesSeeder  =  `insert into branches(address, city_id) values('Ул.Борбади', 1)
					on conflict (address) do nothing`
const PurposesSeeder  =  `insert into purposes(name) values('Заявление'),('Выписка'),('Справка'),('Заказать справку'),('Кредитная карта')
					on conflict (name) do nothing`
const TimesSeeder  =  `insert into times(name) values('9:00'),('9:30'),('10:00'),('10:30'),('11:00'),('11:30	')
					on conflict (name) do nothing`