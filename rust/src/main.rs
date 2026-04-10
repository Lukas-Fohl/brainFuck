use {std::io::Read, std::fs::read};

type TapeSize = u8;

const MAX_SIZE: TapeSize = TapeSize::MAX;

fn execute(tape: &mut [u8;MAX_SIZE as usize], head: &mut TapeSize, pc: &mut usize, program: &Vec<u8>){
    let operation: char = program[*pc] as char;
    match operation{
        '<'=>*head -= 1,
        '>'=>*head += 1,
        '+'=>tape[*head as usize] += 1,
        '-'=>tape[*head as usize] -= 1,
        '.'=>print!("{}", tape[*head as usize] as char),
        ','=>{std::io::stdin().read(&mut tape[*head as usize..*head as usize+1]).unwrap();},
        '['=>{
                if tape[*head as usize] == 0 {
                    let mut depth = 1;
                    while depth > 0 {
                        *pc += 1;
                        if program[*pc] as char == '[' { depth += 1; }
                        if program[*pc] as char == ']' { depth -= 1; }
                    }
                }
            },
        ']'=>{
                if tape[*head as usize] != 0 {
                    let mut depth = 1;
                    while depth > 0 {
                        *pc -= 1;
                        if program[*pc] as char == ']' { depth += 1; }
                        if program[*pc] as char == '[' { depth -= 1; }
                    }
                }
            },
        _  => {} 
    } 
}

fn main() {
    let mut tape: [u8; MAX_SIZE as usize] = [0;MAX_SIZE as usize];
    let mut head: TapeSize = 0;
    let program: Vec<u8> = read("hello.bf").unwrap();
    let mut pc: usize = 0;
    while pc < program.len() {
        execute(&mut tape, &mut head, &mut pc, &program);
        pc += 1;
    }
}
