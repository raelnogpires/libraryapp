DROP DATABASE IF EXISTS LibraryDB;
CREATE DATABASE LibraryDB;

CREATE TABLE LibraryDB.users(
	id INT PRIMARY KEY AUTO_INCREMENT,
  	username VARCHAR(50) NOT NULL,
  	email VARCHAR(70) NOT NULL,
  	`password` VARCHAR(120) NOT NULL
) ENGINE = InnoDB;

CREATE TABLE LibraryDB.authors(
	id INT PRIMARY KEY AUTO_INCREMENT,
  	`name` VARCHAR(50) NOT NULL
) ENGINE = InnoDB;

CREATE TABLE LibraryDB.categories(
	id INT PRIMARY KEY AUTO_INCREMENT,
  	`name` VARCHAR(50) NOT NULL
) ENGINE = InnoDB;

CREATE TABLE LibraryDB.books(
	id INT PRIMARY KEY AUTO_INCREMENT,
  	`name` VARCHAR(70) NOT NULL,
  	description VARCHAR(600),
  	category_id INT NOT NULL,
  	author_id INT NOT NULL,
  	img_url VARCHAR(200),
  	FOREIGN KEY (category_id) REFERENCES LibraryDB.categories (id),
  	FOREIGN KEY (author_id) REFERENCES LibraryDB.authors (id)
) ENGINE = InnoDB;

INSERT INTO LibraryDB.users (username, email, `password`) VALUES
	('admin', 'admin@librarydb.com', 'supersecure72649');

INSERT INTO LibraryDB.authors (`name`) VALUES
	('Clarice Lispector'),
    ('Carlos Drummond de Andrade'),
    ('Hanya Yanagihara'),
    ('Fyodor Dostoevsky'),
    ('Simone de Beauvoir'),
    ('Jean-Paul Sartre');

INSERT INTO LibraryDB.categories (`name`) VALUES
	('Ciências Sociais'),
    ('Ficção Literária'),
    ('Romance'),
    ('Ficção Existencialista'),
    ('Literatura Russa'),
    ('Poesia');
    
INSERT INTO LibraryDB.books (`name`, `description`, category_id, author_id, img_url) VALUES
	('A PAIXÃO SEGUNDO G.H.',
     'A paixão segundo G.H. conta, através de um enredo banal, o pensar e o sentir de G.H., a protagonista-narradora que despede a empregada doméstica e decide fazer uma limpeza geral no quarto de serviço, que ela supõe imundo e repleto de inutilidades.',
     2, 1, 'https://m.media-amazon.com/images/I/51qzgHEl-BL.jpg'),
    ('A DESCOBERTA DO MUNDO',
     'Se nos contos e romances o mistério de uma narrativa envolve o leitor num processo quase que iniciático, nas crônicas esse mistério vai aos poucos sendo desvendado, revelando o mundo pessoal e subjetivo da autora enigmática que viveu no Leme, próximo às areias e ao mar de Copacabana, que tanto apreciava.',
     2, 1, 'https://images-na.ssl-images-amazon.com/images/I/71iAYKeSuIL.jpg'),
    ('Sentimento do mundo',
     'O Drummond de Sentimento do mundo oscila entre diversos polos: cidade x interior, atualidade x memórias, eu x mundo. Perfeita depuração dos livros anteriores, este é um verdadeiro marco.',
     6, 2, 'https://images-na.ssl-images-amazon.com/images/I/41Q6T14Y0EL._SX324_BO1,204,203,200_.jpg'),
    ('Alguma poesia',
     '"Alguma poesia" assinala a estreia de um autor que, então com 28 anos, iria revolucionar a poesia de língua portuguesa no século XX. O livro demonstra já a enorme maturidade do jovem Drummond.',
     6, 2, 'https://images-na.ssl-images-amazon.com/images/I/41DlkipWHvL._SX324_BO1,204,203,200_.jpg'),
    ('Uma vida pequena',
     '"Uma vida pequena" é um dos livros mais surpreendentes, desafiadores, perturbadores e profundamente emocionantes das últimas décadas. Candidato ao Prêmio Pulitzer de Literatura de 2016, além de finalista do Man Booker Prize e do National Book Award.',
     3, 3, 'https://images-na.ssl-images-amazon.com/images/I/519Zz7U8RgL._SX345_BO1,204,203,200_.jpg'),
    ('Crime e Castigo',
     '"Crime e Castigo" é a obra mais célebre de Fyodor Dostoevsky. Neste livro, Raskólnikov, um jovem estudante, pobre e desesperado, perambula pelas ruas de São Petersburgo até cometer um crime que tentará justificar por uma teoria: grandes homens, como César ou Napoleão, foram assassinos absolvidos pela História.',
     5, 4, 'https://images-na.ssl-images-amazon.com/images/I/517DdyXpc5L._SX348_BO1,204,203,200_.jpg'),
    ('Os irmãos Karamázov',
     'Último romance de Dostoevsky, Os irmãos Karamázov representa uma síntese de toda sua produção e é tido por muitos como sua obra-prima. Um marco da literatura universal, influenciou pensadores como Nietzsche e Freud - que o considerava "o maior romance já escrito" - e sucessivas gerações de escritores em todo o mundo.',
     5, 4, 'https://images-na.ssl-images-amazon.com/images/I/51r+M2bjV7L._SX348_BO1,204,203,200_.jpg'),
    ('O segundo sexo',
     'O segundo sexo foi publicado originalmente em 1949 e consagrou Simone de Beauvoir na filosofia mundial. A obra, no entanto, não ficou datada e tornou-se atemporal e definitiva. Simone de Beauvoir analisa a condição da mulher em todas as suas dimensões: sexual, psicológica, social e política e aborda os fatos e os mitos da condição da mulher numa reflexão fascinante.',
     1, 5, 'https://m.media-amazon.com/images/I/41zrzkfl1IL._SY346_.jpg'),
    ('A náusea',
     'A náusea é o primeiro romance de Jean-Paul Sartre, considerado pela crítica e pelo próprio autor o mais perfeito de sua sempre inquieta e inovadora carreira. O protagonista desta história é o intelectual pequeno-burguês Antoine Roquentin, símbolo de uma geração que descobre, horrorizada, a ausência de sentido da vida.',
     4, 6, 'https://images-na.ssl-images-amazon.com/images/I/816WwKmgl5L.jpg');
