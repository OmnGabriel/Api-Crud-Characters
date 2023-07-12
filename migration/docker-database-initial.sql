create table characters(
    id serial primary key,
    name varchar,
    backStory varchar
);

INSERT INTO characters(name, backStory) VALUES
('L (character)', 'L (エル, Eru) is a world-renowned detective who takes on the challenge of catching the mass murderer known as Kira. In his investigation, L becomes suspicious of Light Yagami and makes it his goal to prove that Light is Kira.'),
('Tanjiro Kamado', 'Tanjiro Kamado (竈門炭治郎 Kamado Tanjirō) is the main protagonist of Demon Slayer: Kimetsu no Yaiba. He is a Demon Slayer in the Demon Slayer Corps, who joined to find a remedy to turn his sister, Nezuko Kamado, back into a human and to hunt down and kill demons, and later swore to defeat Muzan Kibutsuji, the King of Demons, in order to prevent others from suffering the same fate as him.');