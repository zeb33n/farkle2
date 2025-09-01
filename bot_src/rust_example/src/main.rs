use anyhow::{Context, Result};
use serde_json;
use std::io::{self, Write};

fn main() -> Result<()> {
    loop {
        let mut buffer = String::new();
        io::stdin().read_line(&mut buffer)?;
        let json: serde_json::Value = serde_json::from_str(buffer.as_str())?;
        if let serde_json::Value::Object(obj) = json {
            if let serde_json::Value::Number(num) =
                obj.get("CurrentScore").context("Score not found")?
            {
                let mut stdout = io::stdout();
                if num.as_u64().context("Not an int")? > 500 {
                    stdout.write_all(b"b")?;
                    stdout.flush()?;
                } else {
                    stdout.write_all(b"r")?;
                    stdout.flush()?;
                }
            }
        }
    }
}
