// export.js
const fs = require('fs');
const path = require('path');

// Ruta al directorio de exportación
const exportPath = path.join(__dirname, 'build');

// Función para mover archivos .html a index.html dentro de su propio directorio
function moveHtmlToIndex(dir) {
    fs.readdirSync(dir).forEach(file => {
        const fullPath = path.join(dir, file);
        const stat = fs.statSync(fullPath);

        if (stat.isDirectory()) {
            moveHtmlToIndex(fullPath);
        } else if (path.extname(file) === '.html' && file !== 'index.html') {
            const newDir = path.join(dir, path.basename(file, '.html'));
            const newIndexPath = path.join(newDir, 'index.html');

            // Crear nuevo directorio y mover index.html
            fs.mkdirSync(newDir, { recursive: true });
            fs.renameSync(fullPath, newIndexPath);
        }
    });
}

// Mover archivos en el directorio de exportación
moveHtmlToIndex(exportPath);
